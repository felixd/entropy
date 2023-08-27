# Entropy

Entropy Playground

## Go lang

Check [go/](go/) directory for Go Lang examples

## Online Services



## Entropy Evaluation

### Data
* Webcam/CCTV Random Number Generator results ~100 MB: https://servertest.online/entropy/20230825-185423-3cdd
* /dev/urandom 4GB (Linux vpn-s 5.4.0-156-generic #173-Ubuntu SMP Tue Jul 11 07:25:22 UTC 2023 x86_64 x86_64 x86_64 GNU/Linux) https://servertest.online/entropy/20230825-134528-9c3c
* macOS https://servertest.online/entropy/20230825-132329-7d78
* Quantum Device (1GB file: https://servertest.online/entropy/20230827-115443-4324
* 

### Software 

* US NIST.gov https://github.com/usnistgov/SP800-90B_EntropyAssessment

## Other things

* https://servertest.online/entropy
* https://wiki.alpinelinux.org/wiki/Entropy_and_randomness
* https://rsmith.home.xs4all.nl/howto/fun-with-encryption-and-randomness.html
* https://webhome.phy.duke.edu/~rgb/General/dieharder.php
* https://medium.com/unitychain/provable-randomness-how-to-test-rngs-55ac6726c5a3


## Linux

### Software

```bash
sudo apt install rng-tools
```

```bash
felixd@vpn-s:~/rng/linux-ent$ cat /dev/random | rngtest --blockcount=100000
rngtest 5
Copyright (c) 2004 by Henrique de Moraes Holschuh
This is free software; see the source for copying conditions.  There is NO warranty; not even for MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.

rngtest: starting FIPS tests...
rngtest: bits received from input: 2000000032
rngtest: FIPS 140-2 successes: 99925
rngtest: FIPS 140-2 failures: 75
rngtest: FIPS 140-2(2001-10-10) Monobit: 10
rngtest: FIPS 140-2(2001-10-10) Poker: 11
rngtest: FIPS 140-2(2001-10-10) Runs: 27
rngtest: FIPS 140-2(2001-10-10) Long run: 28
rngtest: FIPS 140-2(2001-10-10) Continuous run: 0
rngtest: input channel speed: (min=5.116; avg=4125.041; max=19073.486)Mibits/s
rngtest: FIPS tests speed: (min=4.608; avg=146.269; max=150.185)Mibits/s
rngtest: Program run time: 13512568 microseconds
```

### Quantum Random

https://pypi.org/project/quantumrandom/

### Since Kernel 5.18

Since Kernel 5.18 /dev/random and /dev/urandom are the same thing

Finally I have been able to find more explanation on what happened. You can read commit message from Kernel.org explaining situation:

More explanation: https://lwn.net/Articles/884875/

* https://lore.kernel.org/lkml/20220522214457.37108-1-Jason@zx2c4.com/T/#u
* https://git.kernel.org/pub/scm/linux/kernel/git/torvalds/linux.git/commit/?id=ac2ab99072cce553c78f326ea22d72856f570d88
* https://www.theregister.com/2022/03/21/new_linux_kernel_has_improved/
* https://kernelnewbies.org/Linux_5.18#Security

### Monitoring

```bash
[root@server]# echo 1 > /sys/kernel/debug/tracing/events/syscalls/sys_enter_getrandom/enable
[root@server]# cat /sys/kernel/debug/tracing/trace_pipe 
```

## Author

 * Paweł 'felixd' Wojciechowski - FlameIT - [Immersion Cooling](https://flameit.io)
