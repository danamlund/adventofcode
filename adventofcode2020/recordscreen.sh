#!/bin/sh

out=`date "+%Y%m%d_%H%M%S.mp4"`
if [ "$1" != "" ]; then
    out="$1"
fi

ffmpeg -video_size 3839x2160 -f x11grab -i :0.0 -c:v libx265 -crf 28 -preset medium \
       -vf "scale=-2:1080,setpts=PTS*2/30" $out

# 30 fps to 1 fps: 1/30 = 0.0333
# 30 fps to 2 fps: 2/30 = 0.0667
# 30 fps to 5 fps: 5/30 = 0.1667


# convert:
# ffmpeg -i input.mp4 -c:v libx265 -crf 28 -preset medium output.mp4

# speed up
# ffmpeg -i input.mp4 -c:v libx265 -crf 28 -preset medium -vf "setpts=0.2*PTS" output.mp4

# downscale
# ffmpeg -i input.mp4 -c:v libx265 -crf 28 -preset medium -vf "scale=-2:1080" output.mp4

# ffmpeg -i input.mp4 -c:v libx265 -crf 28 -preset medium -vf "scale=-2:720,setpts=0.5*PTS" output.mp4
