# GigaPaste
GigaPaste is the most sophisticated file sharing, url shortener and pastebin all in a single place.

# Why it's better than other similar apps
- works in mobile browsers, can upload file / text with ctrl+v, drag and drop, browse file or through terminal
- Extremely easy to set up, all you need is `go build .` or use the docker-compose.yaml and it's done
- Very easy for modificiations, don't like the style? pick a .css file from [here](https://github.com/dbohdan/classless-css?tab=readme-ov-file#chota) and replace the `static/theme.css`, don't like the layout? the html page is well commented and structured
- Can run on any OS or deployment platforms like repl.it, render, fly.io, etc
- Designed to run efficiently on any system, regardless of its CPU or memory resources
- Can handle gigabytes of file upload and download with fixed memory and cpu usage
- Encryption done right, all your data can be secured with AES & pbkdf2 for passwords
- Decryption is done on the fly, the encrypted data is never decrypted to the disk
- Short & unambiguous URL generation (with letters like ilI1 omitted) with collision detection
- Easy to understand code

# How to build with docker
Just run `docker compose up`  
If you don't want to use docker compose, you need to make a volume for `uploads` and `data` folder (see the docker compose for reference)

# How to build without docker
1. Open terminal
2. type `export CGO_ENABLED=1` (linux) or `SET CGO_ENABLED=1` (windows)
3. `go build .`

# Settings
You can modify the variables inside `data/settings.json`
- fileSizeLimitMB = limit file size (in megabytes)
- textSizeLimitMB = limit text size (in megabytes)
- streamSizeLimitKB = limit file encryption, decryption, upload & download buffer stream size (in kb) to limit memory usage
- streamThrottleMS = add throttle to the encryption, decryption, upload & download buffer to limit cpu usage
- pbkdf2Iterations = key derivation algorithm iteration, the higher the better, but 100000 should be enough
- cmdUploadDefaultDurationMinute = default file duration if you upload file through curl if duration is not specified

You can modify CPU/memory usage with `streamSizeLimitKB * (1000/streamThrottleMS)`, the default setting can handle 40 MB of data on file upload, download, encryption & decryption / second, you can tune this down if needed

# Security
For maximum security, it is recommended to encrypt your file before uploading

# Donate
BTC Network: bc1qpcx3r7pa0uyc957lt84duqexqmvupceyzvyxh6
