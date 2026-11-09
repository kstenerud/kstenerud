### Hi, I'm Karl.

I build the software other developers build on, mostly the hard, low-level kind that has to keep
working when everything around it has failed.

- **[KSCrash](https://github.com/kstenerud/KSCrash)** is the crash reporter at the core of most
  independent iOS crash-reporting SDKs: Sentry, Bugsnag, Datadog, Embrace, bitdrift, Raygun and
  Rollbar all ship it. This code is estimated to have been installed over 100 billion times as
  an app SDK, so it MUST NOT FAIL.
- I've built observability SDKs for several of those vendors: **Sentry's first iOS SDK**,
  **Bugsnag's OpenTelemetry-compatible performance SDK**, **Embrace's sampling profiler**, and
  **bitdrift's crash-resilient storage** (in Rust).

### Now

**[Orpheum](https://github.com/kstenerud/orpheum)** is a self-hosted media server for your own films
and series, which you can watch from anywhere with no VPN and no open ports. Coding agents built it
in seven weeks under my direction, behind tests and gates designed to make their output trustworthy.
[How it's built](https://github.com/kstenerud/orpheum/blob/main/docs/contributors/how-it-is-built.md).

Built along the way:
- **[ffmpeg-fdproxy](https://github.com/kstenerud/ffmpeg-fdproxy)** runs FFmpeg fully confined (no
  network, no file-system path) while it reads untrusted media. Audited spec drafts, fuzzed under
  ASan and UBSan.
- **[yoloAI](https://github.com/kstenerud/yoloai)** gives coding agents a disposable sandbox. The
  agent works on a copy, you review the diff, and credentials never enter the box.

### Also

- **[BONJSON](https://github.com/kstenerud/bonjson)** is a binary drop-in replacement for JSON that
  decodes a median 2.2× faster than yyjson, the fastest JSON library. Implementations in C, Go, Rust and Swift.
- **[Musashi](https://github.com/kstenerud/Musashi)** is a Motorola 680x0 emulator I wrote in the
  '90s. MAME still ships my CPU cores.
- [ObjectAL](https://github.com/kstenerud/ObjectAL-for-iPhone) ·
  [iOS-Universal-Framework](https://github.com/kstenerud/iOS-Universal-Framework) ·
  [KSReachability](https://github.com/kstenerud/KSReachability)

### Work with me

I'm available for contract work and senior roles, remote or around Berlin: crash reporting,
observability SDKs, security-sensitive systems, and making agent-built code production-grade.

- kstenerud@gmail.com
- [LinkedIn](https://www.linkedin.com/in/kstenerud/)
