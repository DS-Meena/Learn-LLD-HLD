In this part, we will learn about high level design of systems.

# Capacity Estimation

A quick and approximate calculation of storage and bandwidth that gives us a further insight.

### Need of capacity Estimation

- How much storage will be needed?
- How many machines will be required?
- What scale is expected from the system?
- Overall, it gives us a rough estimate.

### Cheat sheet

| Number of zeros | Traffic | Storage |
| --- | --- | --- |
| 3 | K | kb |
| 6 | million | mb |
| 9 | billion | GB |
| 12 | trillion | TB |
| 15 | Quadrillion | PB |

Do not spend much time on estimation. Keep assumption values simple. Keep things simple like char (2 bytes), long/double (8 bytes) and Image(300 kb).


## Twitter capacity estimation

Total number of users = 600 million

Daily active users = 200 million

25% of the users tweet, each 2 tweet per day

200 * 25% * 2 = 100 million new tweets per day

200 million * 100 = 20B views per day

### Storage estimation

100 characters per tweet on an average

2 bytes per character

200 bytes + 50 bytes = 250 bytes

Tweet text data = 25GB/day  

1/20 will have an image associated with it: 200KB

1/100 will have a video associated with it: 2MB

Image + Video = (1 + 2) = 3 TB/day 

Total final storage = 25 GB + 3 TB ~= 3 TB/day  


### Bandwidth Estimation

Incoming to server = 3 TB/day ~= 23 MB/sec

Outgoing text tweet data = 20B * 250 bytes ~= 60 MB/sec  

1/20 will have an image associated with it: 200KB 

Outgoing image data ~= 2.5 GB/sec  

1/100 will have a video associated with it: 2MB 

When we show 5 videos, users watch 1 (1/5)

Outgoing video data ~= 1 GB/sec   

Total Outgoing = 60 MB/sec + 2.5 GB/sec + 1 GB/sec ~= 3.5 GB/sec  

This is enough for a rough capacity estimation of Twitter. 