#pragma once

#include <iostream>
#include <map>
#include <string>

#include "Poller/EventPoller.h"
#include "Network/TcpServer.h"
#include "Rtsp/RtspSession.h"
#include "Rtmp/RtmpSession.h"
#include "Http/HttpSession.h"

namespace mediakit {
namespace http {

#define HTTP_FIELD "http."
#define HTTP_PORT 80
const string kPort = HTTP_FIELD"port";
#define HTTPS_PORT 443
const string kSSLPort = HTTP_FIELD"sslport";
onceToken token1([](){
    mINI::Instance()[kPort] = HTTP_PORT;
    mINI::Instance()[kSSLPort] = HTTPS_PORT;
},nullptr);

} // namespace http

namespace rtmp {

#define RTMP_FIELD "rtmp."
#define RTMP_PORT 1935
const string kPort = RTMP_FIELD"port";
onceToken token1([](){
    mINI::Instance()[kPort] = RTMP_PORT;
},nullptr);
} //namespace RTMP

namespace rtsp {

#define RTSP_FIELD "rtsp."
#define RTSP_PORT 554
#define RTSPS_PORT 322
const string kPort = RTSP_FIELD"port";
const string kSSLPort = RTSP_FIELD"sslport";
onceToken token1([](){
    mINI::Instance()[kPort] = RTSP_PORT;
    mINI::Instance()[kSSLPort] = RTSPS_PORT;
},nullptr);
} //namespace Rtsp


} // namespace mediakit