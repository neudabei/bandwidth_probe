# Bandwidth Probe

Test your internet download speed from the command line.
Written in GoLang.

It works by downloading a file from http://ipv4.download.thinkbroadband.com/ and measuring how long it takes.

### Use

Run with  
`$ bandwidth_probe`  
`# Download speed 135.03 Mb/s`  

It will download a file of 20MB in a loop and display the download speed.

You can also specify different file sizes to be downloaded with a command line flag.  
`$ bandwidth_probe -file-size=200`  

You can choose from 5, 20, 200 and 512MB.

### Warning

Running this script consumes real data, so beware if you are on a limited plan. It might incur charges from your internet service provider.
