const express = require('express');
const request = require('request');

const HTML = `<!doctype html>
<html><head>
<meta charset="utf-8">
<style>
body { background: #f7f7f7; font: 14px sans-serif; color: #333; width: 900px; margin: 20px auto; }
pre { margin: 10px 0; padding: 5px 10px; background: #eee; }
a { text-decoration: none; color: azul; }
</style>
</head>
<body><h1>Corset</h1>
<p>Returns any request with <code>Access-Control-Allow-Origin: *</code> headers.</p>
<p><small><a href="https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS">Read more about Cross-Origin Resource Sharing.</a></small></p>
<pre><code>https://corset.herokuapp.com/?url=https://example.com/</code></pre>
<p>Use <a href="https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/encodeURIComponent">encodeURIComponent</a> to encode the URL:</p>
<pre><code>async function myFetch(url) {
    const wrappedUrl = 'https://corset.herokuapp.com/?url=' + encodeURIComponent(url);
    const res = await fetch(wrappedUrl);
    const json = await res.json();
    console.log(json);
    return json;
}
</code></pre>
</body>
</html>
`

const app = express();
app.get('/', (req, res) => {
    if (!req.query.url) {
        res.end(HTML);
        return;
    }
    const url = req.query.url;
    request.get(url, (proxyErr, proxyRes, proxyBody) => {
        res.status(proxyRes.statusCode);
        res.append('Access-Control-Allow-Origin', '*');
        res.end(proxyBody);
    });
});

app.listen(3000, () => console.log('Listening on port 3000'));
