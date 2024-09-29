ARG GO_VERSION=1.22
ARG RUST_VERSION=nightly-2023-12-03
ARG CARGO_CHEF_TAG=0.1.41

FROM ubuntu:20.04

RUN apt-get update --fix-missing && ln -fs /usr/share/zoneinfo/America/New_York /etc/localtime

# Install basic packages
RUN apt-get install build-essential curl wget git pkg-config -y
# Install dev-packages
RUN apt-get update && \
    apt-get install -y --no-install-recommends libclang-dev libssl-dev llvm && \
    rm -rf /var/lib/apt/lists/*

# Install Rust
RUN curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh -s -- -y
ENV PATH="/root/.cargo/bin:${PATH}"
ENV CARGO_HOME=/root/.cargo
# Add Toolchain
ARG RUST_VERSION
RUN rustup toolchain install ${RUST_VERSION}
ARG CARGO_CHEF_TAG
RUN cargo install cargo-chef --locked --version ${CARGO_CHEF_TAG} \
    && rm -rf $CARGO_HOME/registry/

# Install Go
RUN rm -rf /usr/local/go
RUN wget https://raw.githubusercontent.com/anylots/golangTemp/main/go1.22.1.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go1.22.1.linux-amd64.tar.gz
RUN rm go1.22.1.linux-amd64.tar.gz
ENV PATH="/usr/local/go/bin:${PATH}"
