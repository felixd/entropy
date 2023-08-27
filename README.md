# Entropy

Entropy Playground

## Go lang

Check [go/](go/) directory for Go Lang examples

## Entropy Evaluation

### Data
* Webcam/CCTV Random Number Generator results ~100 MB: https://servertest.online/entropy/20230825-185423-3cdd
* /dev/urandom 4GB (Linux vpn-s 5.4.0-156-generic #173-Ubuntu SMP Tue Jul 11 07:25:22 UTC 2023 x86_64 x86_64 x86_64 GNU/Linux) https://servertest.online/entropy/20230825-134528-9c3c
* macOS https://servertest.online/entropy/20230825-132329-7d78

### Software 

* US NIST.gov https://github.com/usnistgov/SP800-90B_EntropyAssessment

## Other things

* https://servertest.online/entropy
* https://wiki.alpinelinux.org/wiki/Entropy_and_randomness
* https://rsmith.home.xs4all.nl/howto/fun-with-encryption-and-randomness.html
* https://webhome.phy.duke.edu/~rgb/General/dieharder.php
* https://medium.com/unitychain/provable-randomness-how-to-test-rngs-55ac6726c5a3


## Linux Kernel

Since Kernel 5.18 /dev/random and /dev/urandom are the same thing

Finally I have been able to find more explanation on what happened. You can read commit message from Kernel.org explaining situation:

More explanation: https://lwn.net/Articles/884875/

* https://lore.kernel.org/lkml/20220522214457.37108-1-Jason@zx2c4.com/T/#u
* https://git.kernel.org/pub/scm/linux/kernel/git/torvalds/linux.git/commit/?id=ac2ab99072cce553c78f326ea22d72856f570d88
* https://www.theregister.com/2022/03/21/new_linux_kernel_has_improved/
* https://kernelnewbies.org/Linux_5.18#Security

## Author

 * Paweł 'felixd' Wojciechowski - FlameIT - [Immersion Cooling](https://flameit.io)
