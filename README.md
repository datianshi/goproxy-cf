goproxy-cf
==========

* Git clone https://github.com/datianshi/goproxy-cf
* cd goproxy-cf
* Config the proxy mapping in a map
E.g.

<pre><code>
&map[string]string{"/": "http://liberty-music.cfapps.io"}
</pre></code>

The above code proxy / to another server : http://liberty-music.cfapps.io

* Push to cloudfoundry
cf push go-proxy -b https://github.com/cloudfoundry/go-buildpack -c "goproxy-cf"

Now access go-proxy.cfapps.io will proxy to http://liberty-music.cfapps.io
  

