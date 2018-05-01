
# Urbano, the [alf.io](https://alf.io) update agent

## what's this?

Urbano is an agent which takes care of updating your [alf.io](https://alf.io) instance to the latest stable release. 
It is automatically included by [Swicket](https://swicket.io) on each "private" deployment of alf.io. 

## Prerequisite

You should have already alf.io installed on a systemd-capable linux. 
Urbano needs run the following commands without being asked for a password:

```
/usr/bin/systemctl stop alfio
/usr/bin/systemctl start alfio
```

to allow that, create a file under /etc/sudoers.d/ with the following content:

```
username ALL=NOPASSWD: /usr/bin/systemctl stop alfio
username ALL=NOPASSWD: /usr/bin/systemctl start alfio
```
replace * username * with the actual username

## Build / Run on your machine

In order to build / run Urbano on your machine, clone this repository under

``` sh
$GOPATH/github.com/swicket/urbano/
```

## Aknowledgements

Urbano uses the great [grab](https://github.com/cavaliercoder/grab) download manager for downloading the latest version of alf.io. Grab is released under BSD-3-Clause license.