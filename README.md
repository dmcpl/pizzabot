# Slice Code Challenge

- [Building and running](#Building-and-running)
- [Parameters](#Parameters)
- [Optional arguments](#Optional-arguments)
- [Future Improvements](#Future-improvements)
- [Problem Introduction](#Introduction)


## Building and running

From the current directory you can build an executable in the bin/ directory then
run the program execute the following commands:

`go build -o bin/pizzabot main.go` then `bin/pizzabot <params>`

You may also run the program directly without producing an executable using the
following:

`go run main.go <params>`

### Parameters

The main game arguments are comprised of a grid size and coordinates of the path
the bot should take around the grid.

Example:
`"5x5 (1, 3) (4, 4) (1, 1) (2, 1) ..."`

In an effort to make this a bit  more user-friendly I pattern match on the whole
line for grid size, accepting the first occurrence ignoring others. I do the same
for each of the path points specified. I ignore whitespace between and around the
parameters. Anything not recognised is disregarded.

Note: this argument should be surrounded by double quotes and should appear last
on the line of command line arguments.

### Optional arguments

`-v`
The default output of the Pizzabot program is a single line of directions and drops
as specified in the challenge. I also included a _verbose_ output mode that would
show the time taken to figure out the path(s), and display all paths attempted
depending on which _algorithm_ (see below) was used.

`-a [algorithm]`
You can specify the type of pathfinding _algorithm_ the Pizzabot will use to traverse the points
specified. 'algorithm' is one of the following:

- `CP` **(Closest Point)**:
This is the default pathfinding algorithm. The next point chosen by the bot will be 
the closest in steps, if there are points equally distant, the first in the list will
be chosen.

- `ordered` **(Ordered)**:
The points are traversed in the order that the user specifies on the command line. This
differs from other algorithms in that it will travel away from and back to a point
that is specified twice or more rather than making two drops there.

- `treeCP` **(Node Tree Closest Point)**:
This algorithm operates similarly to the Closest Point, the main difference is that when
two subsequent point options are equidistant from the current point both paths are stored
and followed. This is achieved by storing path points in Nodes of a decision tree. The default
output will just display the first of the shortest routes, the verbose output will show every
route attempted.

- `treeBF` **(Node Tree Brute Force)**:
Every single combination of remaining path point will be attempted. For example if the path has
seven points, all seven points will be attempted first, then from each of those points the remaining
six and so on. If n was the number of points then n! (factorial) paths will be produced. This can
get out of hand fairly quickly and produces some interesting benchmark times with verbose output.
You will be probably be best piping this output to file.

Example line:
`bin/pizzabot -v -a treeBF "5x5 (1, 3) (4, 4) (1, 1) (2, 1)"`


### Future improvements
- Specifiable origin point
- I/O from file sources
- Show which parts of the command line were unparsed, show currently parsable parameters and ask
for permission to continue.

## Introduction

As part of Slice's commitment to reducing bias in the interview process, we're
asking you to complete a code challenge. The challenge is intended to be
respectful of your time, language- and platform-neutral, appropriate for
engineers of all skill levels, and (hopefully) fun. All challenges are stripped
of identifying information and judged against a rubric by two independent
reviewers. You can make the anonymization process easier for us by not
including your name in source files or documentation.

Slice engineers work predominantly in Ruby, Javascript, Python, Swift, and
Android Kotlin, but you're welcome to complete the challenge in the programming
language of your choice. If we believe we're not qualified to evaluate it,
we'll let you know.

If you successfully complete the challenge and agree to a formal interview,
we may ask you to pair-program with one of our engineers on an extension to
your submission as part of that process.

Please submit the solution to your challenge as a tarball, with clear
instructions on how to execute it.

## Rubric

Here's what we're looking for:

* _Correctness_. Does the code fulfill the requirements of the challenge?
* _Readability_. Is the code well-structured by the standards of the host
  language? Is it simple and clean? Does it reflect the domain of the problem?
* _Fit and polish_. Is there a README? A build script? Are there spelling
  errors or extraneous comments? How does it handle unspecified behavior?
* _Test coverage_. Not every developer writes tests, and that's okay. But we
  do. (Most of the time.)

## Challenge: Pizzabot (also see PDF)

As part of our continuing commitment to the latest cutting-edge pizza
technology research, Slice is working on a robot that delivers pizza. We call
it _(dramatic pause)_: Pizzabot. Your task is to instruct Pizzabot on how to
deliver pizzas to all the houses in a neighborhood.

In more specific terms, given a grid (where each point on the grid is one
house) and a list of points representing houses in need of pizza delivery,
return a list of instructions for getting Pizzabot to those locations and
delivering. An instruction is one of:

```
N: Move north
S: Move south
E: Move east
W: Move west
D: Drop pizza
```

Pizzabot always starts at the origin point, (0, 0). As with a Cartesian
plane, this point lies at the most south-westerly point of the grid.

Therefore, given the following input:

```sh
$ ./pizzabot "5x5 (1, 3) (4, 4)"
```

one correct solution would be:

```
ENNNDEEEND
```

In other words: move east once and north thrice; drop a pizza; move east thrice
and north once; drop a final pizza.

If you'd prefer to avoid stdin, or work predominantly in a platform that makes
it difficult to use, the equivalent solution expressed as an integration test is
just fine. The API is entirely up to you, as long as the test exercises
functionality that accepts and returns properly formatted strings:

```
assertEqual(pizzabot("5x5 (1, 3) (4, 4)"), "ENNNDEEEND")
```

There are multiple correct ways to navigate between locations. We do not take
optimality of route into account when grading: all correct solutions are good
solutions.

To complete the challenge, please solve for the following _exact input_:

```sh
5x5 (0, 0) (1, 3) (4, 4) (4, 2) (4, 2) (0, 1) (3, 2) (2, 3) (4, 1)
```

Keep it simple, and have fun!

