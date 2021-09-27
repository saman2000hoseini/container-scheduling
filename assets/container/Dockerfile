FROM ubuntu

ENV TZ=Asia/Tehran \
    PATH="/app:${PATH}"

RUN apt update

RUN apt install -y \
      gcc \
      musl-dev \
      build-essential \
      python \
      python3-pip

WORKDIR /app

COPY . .

RUN g++ -o min min.cpp
RUN g++ -o max max.cpp
RUN g++ -o average average.cpp
RUN g++ -o sort sort.cpp
RUN g++ -o word_count word_count.cpp

RUN rm *.cpp

RUN pip install pigar

RUN echo 'alias syncdependencies="echo y | pigar -P ./temp"' >> ~/.bashrc

RUN printf "#!/bin/bash\n echo y | pigar -P ./temp" > /app/syncdependencies.sh && chmod +x /app/syncdependencies.sh

ENTRYPOINT ["/bin/bash"]