# cli-create
A tool to create c++ program with boilerplate

```
$ ./cli-create main
$ main.cpp created
$ cat main.cpp

#include <bits/stdc++.h>
#include <chrono>

#define TIMER_START auto TIME_START = std::chrono::high_resolution_clock::now()
#define TIMER_END auto TIME_END = std::chrono::high_resolution_clock::now()
#define TIMECHECK std::cout << std::chrono::duration_cast<std::chrono::milliseconds>(TIME_END - TIME_START).count() << "ms" << std::endl

typedef double f64;
typedef long long i64;
typedef int i32;
typedef pair<i32, i32> pi32;
typedef unsigned long long u64;
typedef unsigned int u32;
typedef vector<i32> vi32;
typedef deque<i32> di32;

int main() {
    const int N = 1000000;

    TIMER_START;

    // Simulating some time-consuming operation, like a loop
    for (int i = 0; i < 1000000; ++i) {

        // Perform some computation
        int result = i * i;
    }

    TIMER_END;
    TIMECHECK;

    return 0;
}
```


