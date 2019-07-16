/*
 * Copyright 2017-2019 Baidu Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package com.baidu.openrasp.plugin.js;

import java.util.*;
import com.baidu.openrasp.request.AbstractRequest;
import com.jsoniter.output.JsonStream;
import com.baidu.openrasp.v8.ByteArrayOutputStream;

public class Context extends com.baidu.openrasp.v8.Context {
    
    public AbstractRequest request = null;

    public static void setKeys() {
        setStringKeys(new String[] { "path", "method", "url", "querystring", "protocol", "remoteAddr", "appBasePath", "requestId" });
        setObjectKeys(new String[] { "json", "server", "parameter", "header" });
        setBufferKeys(new String[] { "body" });
    }

    public Context(AbstractRequest request) {
        this.request = request;
    }

    public String getString(String key) {
        if (key.equals("path"))
            return getPath();
        if (key.equals("method"))
            return getMethod();
        if (key.equals("url"))
            return getUrl();
        if (key.equals("querystring"))
            return getQuerystring();
        if (key.equals("appBasePath"))
            return getAppBasePath();
        if (key.equals("protocol"))
            return getProtocol();
        if (key.equals("remoteAddr"))
            return getRemoteAddr();
        if (key.equals("requestId"))
            return getRequestId();
        return null;
    }

    public byte[] getObject(String key) {
        if (key.equals("body"))
            return getBody();
        if (key.equals("json"))
            return getJson();
        if (key.equals("header"))
            return getHeader();
        if (key.equals("parameter"))
            return getParameter();
        if (key.equals("server"))
            return getServer();
        return null;
    }

    public byte[] getBuffer(String key) {
        if (key.equals("body"))
            return getBody();
        return null;
    }

    public String getPath() {
        try {
            return request.getRequestURI();
        } catch (Exception e) {
            return null;
        }
    }

    public String getMethod() {
        try {
            return request.getMethod().toLowerCase();
        } catch (Exception e) {
            return null;
        }
    }

    public String getUrl() {
        try {
            return request.getRequestURL().toString();
        } catch (Exception e) {
            return null;
        }
    }

    public String getQuerystring() {
        try {
            return request.getQueryString();
        } catch (Exception e) {
            return null;
        }
    }

    public String getAppBasePath() {
        try {
            return request.getAppBasePath();
        } catch (Exception e) {
            return null;
        }
    }

    public String getProtocol() {
        try {
            return request.getProtocol().toLowerCase();
        } catch (Exception e) {
            return null;
        }
    }

    public String getRemoteAddr() {
        try {
            return request.getRemoteAddr();
        } catch (Exception e) {
            return null;
        }
    }

    public String getRequestId() {
        try {
            return request.getRequestId();
        } catch (Exception e) {
            return null;
        }
    }

    public byte[] getBody() {
        try {
            return request.getBody();
        } catch (Exception e) {
            return null;
        }
    }

    public byte[] getJson() {
        try {
            String contentType = request.getContentType();
            if (contentType != null && contentType.contains("application/json")) {
                return getBody();
            }
            return null;
        } catch (Exception e) {
            return null;
        }
    }

    public byte[] getHeader() {
        try {
            Enumeration<String> headerNames = request.getHeaderNames();
            if (headerNames == null || !headerNames.hasMoreElements()) {
                return null;
            }
            HashMap<String, String> headers = new HashMap<String, String>();
            while (headerNames.hasMoreElements()) {
                String key = headerNames.nextElement();
                String value = request.getHeader(key);
                headers.put(key.toLowerCase(), value);
            }
            ByteArrayOutputStream out = new ByteArrayOutputStream();
            JsonStream.serialize(headers, out);
            out.write(0);
            return out.getByteArray();
        } catch (Exception e) {
            return null;
        }
    }

    public byte[] getParameter() {
        try {
            Map<String, String[]> parameters = request.getParameterMap();
            if (parameters == null || parameters.isEmpty()) {
                return null;
            }
            ByteArrayOutputStream out = new ByteArrayOutputStream();
            JsonStream.serialize(parameters, out);
            out.write(0);
            return out.getByteArray();
        } catch (Exception e) {
            return null;
        }
    }

    public byte[] getServer() {
        try {
            Map<String, String> server = request.getServerContext();
            ByteArrayOutputStream out = new ByteArrayOutputStream();
            JsonStream.serialize(server, out);
            out.write(0);
            return out.getByteArray();
        } catch (Exception e) {
            return null;
        }
    }
}