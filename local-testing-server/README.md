# Local https server used for hot reloading the extension 
Serves files required by the Twitch extensions **Local Test** feature


## Usage

### 1. Create locally-trusted development certificate
```
mkcert -install localhost
```
Running this command will generate two files `localhost-key.pem` and `localhost.pem`

### 2. Move certificate files in this project directory
Make sure that both certificate files are in the same directory as `server.py`

### 3. Run the server
```
python server.py <path_to_extension>
```

### Example

Running from `local-testing-server` directory

```
mkcert -install localhost
python server.py "../extension/src"
```
