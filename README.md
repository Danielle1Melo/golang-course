<p align="center">
  <img src="/assets/gopher-dance-long.gif">
</p>


# Learning GoLang

This repository contains my studies and implementations in Go related to **finite fields** and **elliptic curves**, focusing on the fundamentals used in Bitcoin.  
I am currently following a GoLang course and applying the concepts learned in the context of mathematics applied to cryptography.

## What is a finite field?

A **finite field** is a limited set of numbers where we can perform addition, subtraction, multiplication and division operations (except for zero), and the result of any operation **always remains within the set**.

In the case of Bitcoin, we use a finite field defined by a very large prime number `p`.  
This means that any operation performed is always reduced **mod p**, ensuring that the results stay within the field boundaries.

