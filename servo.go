/*
Copyright 2021 Manish Sahani (kalkayan)

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package servo

import "fmt"

/*

     _______. _______ .______     ____    ____  ______
    /       ||   ____||   _  \    \   \  /   / /  __  \
   |   (----`|  |__   |  |_)  |    \   \/   / |  |  |  |
    \   \    |   __|  |      /      \      /  |  |  |  |
.----)   |   |  |____ |  |\  \----.  \    /   |  `--'  |
|_______/    |_______|| _| `._____|   \__/     \______/


 Servo is a REST-based go server framework (micro) for creating high-performance
 API applications. It's a good starting point for developers, allowing them to
 focus on creating something amazing while the servo sweat over the smaller
 details.
*/

// Hello returns a greeting for the named person.
func Hello(name string) string {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
