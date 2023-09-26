# Goroutine Visualization

This Go program provides a visualization of the differences between running code with and without goroutines. It uses the `svgo` library to generate SVG images that illustrate the timeline of two processes.

## Overview

The code creates a web server that serves an SVG image at the "/visualize" endpoint. The image consists of two timelines, one for each process. The processes are simulated, and their progress is visually represented on the timeline.

- **Process 1 Timeline**: Shown in red, represents a process without goroutines.
- **Process 2 Timeline**: Shown in green, represents a process with goroutines.


## Differences

Placeholder images illustrating the differences between the two processes are shown below:

### Process 1 (Without Goroutines)
![Process 1 (No Goroutines)](https://github.com/oguzhantasimaz/Go-Concurrency-Examples-In-Detail/blob/main/VisualizeConcurrencyWithSVG/without_concurrency.png)

### Process 2 (With Goroutines)
![Process 2 (With Goroutines)](https://github.com/oguzhantasimaz/Go-Concurrency-Examples-In-Detail/blob/main/VisualizeConcurrencyWithSVG/with_concurrency.png)
