This project attempts to bring frequency domain math to Golang. These routines will eventually be highly optimized for quick transforms of large arrays.

Below is the pseudo code followed from this wikipedia article: https://en.wikipedia.org/wiki/Cooley–Tukey_FFT_algorithm
```
X0,...,N−1 ← ditfft2(x, N, s):             DFT of (x0, xs, x2s, ..., x(N-1)s):
    if N = 1 then
        X0 ← x0                                      trivial size-1 DFT base case
    else
        X0,...,N/2−1 ← ditfft2(x, N/2, 2s)             DFT of (x0, x2s, x4s, ...)
        XN/2,...,N−1 ← ditfft2(x+s, N/2, 2s)           DFT of (xs, xs+2s, xs+4s, ...)
        for k = 0 to N/2−1                           combine DFTs of two halves into full DFT:
            t ← Xk
            Xk ← t + exp(−2πi k/N) Xk+N/2
            Xk+N/2 ← t − exp(−2πi k/N) Xk+N/2
        endfor
    endif
```

Example from numpy:
input = [np.pi/4, np.pi/2, 3*np.pi/4, np.pi]

output = array( [ 7.85398163+0.j,
                -1.57079633+1.57079633j,
                -1.57079633+0.j,
                -1.57079633-1.57079633j])

output from this routine = (7.853981633974483+0i)
                           (-1.5707963267948966+1.5707963267948966i)
                           (-1.5707963267948966+0i)
                           (-1.5707963267948966-1.5707963267948966i)
