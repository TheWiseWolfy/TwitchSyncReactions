from http.server import SimpleHTTPRequestHandler, HTTPServer
import ssl
import os
import sys

current = os.getcwd()
os.chdir(sys.argv[1])

cert_file = os.path.join(current, "localhost.pem")
key_file = os.path.join(current, "localhost-key.pem")

class RequestHandler(SimpleHTTPRequestHandler):
    def end_headers(self):
        self.send_header("Cache-Control", "no-cache, no-store, must-revalidate")
        self.send_header("Pragma", "no-cache")
        self.send_header("Expires", "0")
        SimpleHTTPRequestHandler.end_headers(self)

httpd = HTTPServer(
    ('', 8080),
    RequestHandler
)

sslctx = ssl.SSLContext(ssl.PROTOCOL_TLS_SERVER)
sslctx.check_hostname = False
sslctx.load_cert_chain(certfile=cert_file, keyfile=key_file)
httpd.socket = sslctx.wrap_socket(httpd.socket, server_side=True)
httpd.serve_forever()
