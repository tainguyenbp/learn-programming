#!/bin/bash

# Function to parse Dockerfile and extract ARG variable names
parse_dockerfile() {
    arg_names=()
    while IFS= read -r line; do
        # if [[ $line =~ ^ARG\ (.+)$ ]]; then # get ARG all not exclude
        if [[ $line =~ ^ARG\ ([^\=]+)$ ]]; then # not get ARG env have = character

            arg_names+=("${BASH_REMATCH[1]}")
        fi
    done < Dockerfile
}

# Function to construct the docker build command with --build-arg flags
construct_build_command() {
    build_command=("docker" "build")
    added_args=()  # Array to keep track of added arguments
    for arg_name in "${arg_names[@]}"; do
        # Check if the argument has already been added
        if [[ ! " ${added_args[@]} " =~ " ${arg_name} " ]]; then
            build_command+=("--build-arg" "${arg_name}=\${$arg_name}")
            added_args+=("${arg_name}")
        fi
    done
    build_command+=(".")
}

# Main function
main() {
    parse_dockerfile
    construct_build_command
    echo "${build_command[@]}"
    # Uncomment the below line to execute the docker build command
    # eval "${build_command[@]}"
}
main
