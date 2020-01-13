#ifndef TEST_H 
#define TEST_H 
#include <stdio> 
#ifndef MATH_H
#define MATH_H

int add(int a, int b);
#endif
#ifndef TEST_IMPLEMENTATION
#define TEST_IMPLEMENTATION

int add(int a, int b) {
  return a + b;
}
#endif
#endif
