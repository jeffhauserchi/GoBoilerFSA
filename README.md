# GoBoilerFSA
Fullstack Academy's Boilermaker adapted for Golang backend -- stackathon/personal enrichment project

This is a work-in-progress, looking at swapping the node.js backend of https://github.com/FullstackAcademy/boilermaker
with one written in Go.

The goal is to see if this can be done without making changes to the frontend build.

for express.js:
evaluated:  https://github.com/DronRathore/goexpress, https://github.com/kataras/iris
selected: iris -- worked well out-of-the-box

for sequelize orm:
evaluated: https://github.com/jinzhu/gorm, https://github.com/go-xorm/xorm
selected: gorm -- table creation most closely matched sequelize

for socket.io:
evaluated:  iris-included ws api, https://github.com/googollee/go-socket.io
selected:  evaluation continuing, go-socket.io is able to establish connection but not matching up with js socket.io
in boilermaker.

for passport:
pending
