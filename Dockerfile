FROM arm32v7/golang:1.10 as gobuild
RUN mkdir -p /go/src/github.com/paperplace/
WORKDIR /go/src/github.com/paperplace/
ADD encodeit.go .
ADD qr /go/src/rsc.io/qr
RUN go build github.com/paperplace
FROM arm32v7/python
ADD epaper/2.7inch_e-paper/raspberrypi/python .
RUN pip install spidev
RUN pip install rpi.gpio
RUN pip install pillow
RUN sed -i -e 's/\(self.width \* self.height \/ 8\)/int\(self.width \* self.height \/ 8\)/' -e  's/(x + y \* self.width) \/ 8/int(&)/' epd2in7.py 
RUN sed -i 37,53d main.py
CMD python main.py
