# SPDX-License-Identifier: Apache-2.0
# Copyright Authors of Kerria

# This Makefile help generation tool finds markers (`##`) and converts them into parts of help output.
#
# `##` at the start of a line denotes a section, and the text 1 (one) space after is the section break text.
# Example:
#   `## Section Break Text`
#
# `##` after a `.PHONY` target denotes the description for that target.
# It assumes there is only 1 (one) `.PHONY` target specified.
# Example:
#   `.PHONY: some-target  ## Target description`
#
# `##` after a target denotes the description for that target.
# Example:
#   `some-target: dep1 dep2 dep3  ## Target description`

BEGIN {
    SectionColor = "\033[1;34m"
    TargetColor = "\033[36m"
    ResetColor = "\033[0m"
}

{
    # Section markers
    if ($0 ~ /^## /) {
        title = substr($0, 4)  # removes "## "
        printf "\n%s%s%s\n", SectionColor, title, ResetColor
    }
    # .PHONY markers
    else if ($0 ~ /^\.PHONY:.*## /) {
        split($0, parts, /## /)
        target = substr(parts[1], 9)  # removes ".PHONY: "
        description = parts[2]
        printf "%s%-30s%s %s\n", TargetColor, target, ResetColor, description
    }
    # Target markers
    else if ($0 ~ /^[a-zA-Z_-]+:.*## /) {
        split($0, parts, /:.*## /)
        target = parts[1]
        description = parts[2]
        printf "%s%-30s%s %s\n", TargetColor, target, ResetColor, description
    }
}
