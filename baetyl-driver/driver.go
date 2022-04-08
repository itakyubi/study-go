package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	dm "github.com/baetyl/baetyl-go/v2/dmcontext"
	"github.com/baetyl/baetyl-go/v2/errors"
	"github.com/baetyl/baetyl-go/v2/log"
	v1 "github.com/baetyl/baetyl-go/v2/spec/v1"
	"gopkg.in/yaml.v2"
)

type driver struct {
	ctx        dm.Context
	devConfigs map[string]DeviceConfig
	log        *log.Logger
}

// 解析各个子设备配置
func newDriver(ctx dm.Context) (*driver, error) {
	devConfigs := make(map[string]DeviceConfig)
	for _, dev := range ctx.GetAllDevices() {
		var devConfig DeviceConfig
		access, err := ctx.GetDeviceAccessConfig(&dev)
		if err != nil {
			return nil, err
		}
		if err := yaml.Unmarshal([]byte(*access.Custom), &devConfig); err != nil {
			return nil, err
		}
		props, err := ctx.GetDevicePropertiesConfig(&dev)
		if err != nil {
			return nil, err
		}
		var propCfgs []Property
		for _, prop := range props {
			cfg := Property{Name: prop.Name}
			if err := yaml.Unmarshal([]byte(*prop.Visitor.Custom), &cfg); err != nil {
				return nil, err
			}
			propCfgs = append(propCfgs, cfg)
		}
		devConfig.Properties = propCfgs
		devConfigs[dev.Name] = devConfig
		// ensure device is ready, then tell cloud all the devices is online
		if err := ctx.Online(&dev); err != nil {
			return nil, err
		}
	}
	d := &driver{
		ctx:        ctx,
		devConfigs: devConfigs,
		log:        ctx.Log().With(log.Any("module", "custom driver")),
	}
	// 注册数据通知和事件通知回调函数
	if err := ctx.RegisterDeltaCallback(d.DeltaCallback); err != nil {
		return nil, err
	}
	if err := ctx.RegisterEventCallback(d.EventCallback); err != nil {
		return nil, err
	}
	return d, nil
}

func (d *driver) start() {
	for _, dev := range d.ctx.GetAllDevices() {
		go d.running(&dev)
	}
}

func (d *driver) stop() {
}

// 针对多个子设备开始定时读数据
func (d *driver) running(dev *dm.DeviceInfo) {
	cfg, ok := d.devConfigs[dev.Name]
	if !ok {
		d.log.Error("device config not exist", log.Any("device", cfg))
		return
	}
	ticker := time.NewTicker(cfg.Interval)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			err := d.execute(dev)
			if err != nil {
				d.log.Error("failed to execute", log.Error(err))
			}
		case <-d.ctx.WaitChan():
			d.log.Warn("task of device stopped", log.Any("device", dev))
			return
		}
	}
}

// 读取数据后通过上报函数发送数据至baetyl-broker
func (d *driver) execute(dev *dm.DeviceInfo) error {
	cfg, _ := d.devConfigs[dev.Name]
	r := v1.Report{}
	for _, p := range cfg.Properties {
		val, err := d.read(&cfg, &p)
		if err != nil {
			return err
		}
		r[p.Name] = val
	}
	if err := d.ctx.ReportDeviceProperties(dev, r); err != nil {
		return err
	}
	return nil
}

// 读取子设备数据
func (d *driver) read(dev *DeviceConfig, prop *Property) (interface{}, error) {
	resp, err := http.Get(dev.Address + prop.Path)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	if val, ok := result[prop.Name]; ok {
		return val, nil
	} else {
		return nil, fmt.Errorf("failed to get property: %s", prop.Name)
	}
}

// 数据通知回调函数，标识云端写入数据至子设备
func (d *driver) DeltaCallback(dev *dm.DeviceInfo, prop v1.Delta) error {
	devConfig, _ := d.devConfigs[dev.Name]
	for _, cfg := range devConfig.Properties {
		for k, v := range prop {
			if k == cfg.Name {
				pld, err := json.Marshal(map[string]interface{}{cfg.Name: v})
				if err != nil {
					return err
				}
				req, err := http.NewRequest("PUT", devConfig.Address+cfg.Path, bytes.NewBuffer(pld))
				if err != nil {
					return err
				}
				_, err = http.DefaultClient.Do(req)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

// 事件回调函数
func (d *driver) EventCallback(dev *dm.DeviceInfo, event *dm.Event) error {
	switch event.Type {
	// 即时上报事件
	case dm.TypeReportEvent:
		if err := d.execute(dev); err != nil {
			return err
		}
	default:
		return errors.New("event type not supported yet")
	}
	return nil
}
