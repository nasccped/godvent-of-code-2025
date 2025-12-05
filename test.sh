#!/bin/sh

usage_tip() {
	echo -e "usage tip: \e[92m./test.sh \e[96m<OPTIONAL_PACKAGE_NUMBER>\e[0m\n";
	echo -e "when no package number is provide, runs \`\e[92mgo test ./...\e[0m\`"
}

TARGET_PACKAGE=""

case "$#" in
	0)
		TARGET_PACKAGE="./..."
		;;
	1)
		TARGET_PACKAGE=$(printf "./solves/day%02d" "$1")
		;;
	*)
		echo -ne "\e[91merror:\e[0m cant run tests with provided args:";
		for i in "$@"; do
			echo -ne " \e[96m$i\e[0m";
		done
		echo -e "\n";
		usage_tip;
		exit 1;
		;;
esac

if [[ "$TARGET_PACKAGE" != "./..." && ! -d "$TARGET_PACKAGE" ]]; then
	echo -ne "\e[91merror:\e[0m provided package doesn't exists (";
	echo -e "\e[96m$TARGET_PACKAGE\e[0m)";
	exit 1;
fi

go test "$TARGET_PACKAGE"
