#pragma once

#ifndef VIDEO_STREAM_H
#define VIDEO_STREAM_H

#include <iostream>
#include <string>
#include <algorithm>
#include <vector>
#include <functional>
#include <queue>
#include <map>
#include <unordered_map>
#include <mutex>
#include <thread>
#include <cstdio>

#include "Thread/WorkThreadPool.h"
#include "Poller/EventPoller.h"
#include "Common/config.h"
#include "mk_media.h"
#include "block_queue.hpp"

#ifdef __cplusplus
extern "C" {
#endif

#include "libavcodec/avcodec.h"
#include "libavformat/avformat.h"
#include "libavutil/avutil.h"  
#include "libswscale/swscale.h"  
#include <libavutil/imgutils.h>  

#ifdef __cplusplus
}
#endif

class VideoStream {
public:
    void init();
    void receive(std::string msg);
    
private:
    bool initMediaServer();

    BlockQueue<std::string> _inputQueue;

    // 解码线程
    std::thread _decoderThread;
    void decoderThreadMain();
    void decode(std::string msg);
    void startMedia(std::string streadKey, int width, int height);
    // yuv dev channel
    std::unordered_map<std::string, mk_media> _mDev;
    std::unordered_map<std::string, std::vector<int>> _resolutions;
    int _fps = 10;

    // 编码线程
    std::thread _encoderThread;
    void encoderThreadMain();
    void inputYuv(AVFrame *pFrame, std::string streamid);

    // codec
    bool _isCodecReady;
    bool initCodec();
    AVCodec * _pCodec = nullptr;
    AVCodecContext *_pContex = nullptr;
    BlockQueue<AVFrame*> _decoderFrames;
    BlockQueue<AVFrame*> _encoderFrames;
    std::map<std::string, toolkit::SmoothTicker> _aTicker;

    // file
    int readFile(std::string fileName);
    const static int _fileReadBufSize = 1024*1024;
    unsigned char _fileReadBuf[_fileReadBufSize];

};
#endif