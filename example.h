#ifndef EXAMPLE_H 
#define EXAMPLE_H 
#include <stdio> 
#ifndef MATH_H
#define MATH_H

int add(int a, int b);
#endif
#ifndef EXAMPLE_IMPLEMENTATION
#define EXAMPLE_IMPLEMENTATION

int add(int a, int b) {
  return a + b;
}
#endif
#endif
