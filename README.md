# Go Gpx Editor

If you, like me, occasionally forget the Garmin turned on at the end of the lap, or turn it on too soon, this simple solution could help you. It's a personal project with the aim of generating a new GPX file starting from a source file, mainly developed to test some technologies I would like to improve.

This GPX editor is wrote in Golang. The Cobra package is used to provide a simple CLI interface to directly modify the GPX file.

This solution is `work in progress`, there are a lot of commands I would like to integrate here.

I tested it with file GPX extracted from Garmin Connect, in Linux O.S.  

## How does it work?

Extract the Garmin GPX file to modify, get the path of the file, and run one of the commands listed below. At the end of the process, a new file GPX will be create, able to be upload. 

## Commands implemented

- **keepuntil**: keeps unmodified the GPX trace until the end time setted via CLI, the rest of the lap is cutted. 
`go run main.go keepuntil -p=/media/user/DATA/activity.gpx -t=1h6m0s`

- **keepfrom**: keeps unmodified the GPX trace from the start time setted via CLI, the rest of the lap is cutted. 
`go run main.go keepfrom -p=/media/user/DATA/activity.gpx -t=0h6m0s`

- **merge**: combines several GPX tracks, and produces one single track.
`go run main.go merge --path=/media/user/DATA/activity_1.gpx --path=/media/user/DATA/activity_2.gpx`

## Commands in progress

- `extractgpx`

## Next steps

- Docker containerization and gRPC integration. 