#!/usr/bin/env bash

set -eu

pkg_name="easy-fz-chromeos"
pkg_version="0.0.2"

os_list=("windows" "darwin" "linux")
arch_list=("amd64" "arm64")

# 圧縮するか
compression=false
compression_format="tar.xz"

# ディレクトリが存在するか確認
if [[ ! -d bin ]]; then
	mkdir -p bin
fi

# OSに合わせてビルド
for os in "${os_list[@]}"; do
	for arch in "${arch_list[@]}"; do
		output_dir="${pkg_name}_${pkg_version}_${os}_${arch}"

		if [[ ! -d bin/"${output_dir}" ]]; then
			mkdir -p bin/"${output_dir}"
		fi

		cp LICENSE bin/"${output_dir}"
		cp README.md bin/"${output_dir}"

		if [[ "${os}" = "windows" ]]; then
			GOOS=${os} GOARCH=${arch} go build -o bin/"${output_dir}"/"${pkg_name}".exe
		else
			GOOS=${os} GOARCH=${arch} go build -o bin/"${output_dir}"/"${pkg_name}"
		fi
	done

	echo "${os}"版のコンパイルが完了
done

# compressionがtrueの場合圧縮する
if [[ "${compression}" = "true" ]]; then
	for os in "${os_list[@]}"; do
		for arch in "${arch_list[@]}"; do
			output_dir="${pkg_name}_${pkg_version}_${os}_${arch}"
			tar -Jcf bin/"${output_dir}"."${compression_format}" bin/"${output_dir}"
			echo "${output_dir}"フォルダを"${compression_format}"で圧縮
		done
	done
fi
