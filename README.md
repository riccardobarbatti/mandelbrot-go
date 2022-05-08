
# Mandelbrot

The term Mandelbrot set is used to refer both to a general class of fractal sets and to a particular instance of such a set. In general, a Mandelbrot set marks the set of points in the complex plane such that the corresponding Julia set is connected and not computable.

"The" Mandelbrot set is the set obtained from the quadratic recurrence equation

<p align="center">
  <img src="formula.svg" width="100px" alt="mandelbrot" />
</p>

Learn more:

https://en.wikipedia.org/wiki/Mandelbrot_set

You start with a specific number, and then you apply the same recursive formula, again and again, to get the next number in the sequence.

Let’s take the recursive formula 


## Demo
<p align="center">
  <img src="fractal.png" width="1024" height="1024" alt="fractal" /></a>
</p>




## Deployment

To generate fractal image 

```bash
  go run mandelbrot.go > fractal.png
```
