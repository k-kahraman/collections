/*
 * MIT License
 *
 * Copyright (c) 2022 Kaan Kahraman
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package main

import (
	"collections/linear/stack"
	"collections/linear/vector"
	"fmt"
)

func main() {
	vec := vector.New(10, 20, 30)
	stck := stack.New(-1)

	for num := 0; num < 100; num++ {
		stck.Push(num)
	}

	for index := 0; index < 3; index++ {
		elem, _ := vec.Get(index)
		fmt.Printf("Vector element at index %d is %d\n", index, elem)
	}

	for num := 0; num < 102; num++ {
		stckElem, err := stck.Pop()
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Printf("Stack element at top is %d\n", stckElem)
	}

}
