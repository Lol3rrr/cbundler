#ifndef SHL_H 
#define SHL_H 
#include <stdio> 
#ifndef MATH_H
#define MATH_H

int add(int a, int b);
#endif
#ifndef SHL_IMPLEMENTATION
#define SHL_IMPLEMENTATION

int add(int a, int b) {
  return a + b;
}
#endif
#endif
