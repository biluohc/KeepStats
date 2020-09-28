app=keepstats
version=v0.1
commit_id=$(shell git log -1 --pretty="%H"|cut -c1-6)
# src_path = $(shell pwd)
src_path_docker = /opt

# debug or release
# mode=release
ifneq ($(mode), debug)
    target = --release
else
    target =
endif

docker_file_rust = Dockerfile.rust
docker_file_prod = Dockerfile.prod

image_rust = ${app}-rust
image_prod = ${app}-prod

$(info mode: $(target)) # println

$(shell mkdir -p ${HOME}/.cargo/{git,registry})
$(shell touch ${HOME}/.cargo/config)

b:
	cargo build ${target}

image-rust:
	@if [ `docker images | grep ${image_rust} | wc -l` -eq 0 ]; then \
		echo "build rust docker image "; \
        docker build -t ${image_rust}:latest -f ${docker_file_rust} .; \
    else \
        echo "docker image ${image_rust} already exist!";\
    fi

image: build
	@echo "build docker image"; \
	docker build -t ${image_prod}:latest -f ${docker_file_prod} . && \
	docker tag ${image_prod}:latest ${image_prod}:${version}-${commit_id} && \
	echo ${image_prod}:${version}-${commit_id}

build: image-rust
	@echo "docker build ${mode}"; \
	docker run -i --rm \
		-v ${HOME}/.cargo/git:/root/.cargo/git \
		-v ${HOME}/.cargo/config:/root/.cargo/config \
	    -v ${HOME}/.cargo/registry:/root/.cargo/registry \
	    -v `pwd`:$(src_path_docker) \
	    --workdir $(src_path_docker) \
		--network host \
	    ${image_rust}:latest \
	    bash -c "cd $(src_path_docker)/ && cargo build ${target}"

run:
	docker run -d --restart always --network host -v ${PWD}:/opt --name ${app} ${image_prod}

clean:
	rm -fr target/*

