# SPDX-License-Identifier: Apache-2.0
# Copyright Authors of Kerria

# This Makefile help generation tool finds the requested (`-v request=`) make target and converts it into extended help
# output. It detects description markers (`##`) similarly to `make-help.awk` and detects extended description markers
# (`@##` after a tab character) in the recipe for the requested target. This information is used to generate the help
# output.
#
# Example:
#   .PHONY: some-target  ## Target description
#   some-target: dep1 dep2 dep3
#       @## Extended description
#       @## Extended description continued...
#       recipe command
#       @# Comment not part of extended description
#       @## Extended description continued...

BEGIN {
    SectionColor = "\033[1;34m"
    TargetColor = "\033[36m"
    ErrorColor = "\033[1;31m"
    ResetColor = "\033[0m"

    Request = substr(request, 6)  # removes "help-"
    TargetLatch = 0
    TargetFound = 0
    ExtendedDescriptionExists = 0
    TargetPattern = "^" Request ":"

    Description = ""
    PhonyDescriptionPattern = "^\\.PHONY:\\s+" Request "\\s+## "
    TargetDescriptionPattern = "^" Request ":.*## "
}

{
    # Get short description
    # .PHONY markers
    if ($0 ~ PhonyDescriptionPattern) {
        split($0, parts, /## /)
        Description = parts[2]
    }
    # Target markers
    else if ($0 ~ TargetDescriptionPattern) {
        split($0, parts, /## /)
        Description = parts[2]
    }

    # If target latch is enabled, the current context is in the requested target, print long description
    if (TargetLatch && $0 ~ /^\t@## /) {
        ExtendedDescriptionExists = 1
        split($0, parts, /@## /)
        printf "  %s\n", parts[2]
    }
    # Disable the latch when the current target recipe has ended
    else if (TargetLatch && $0 !~ /^\t|^$/) {
        TargetLatch = 0
    }

    # If target matches request, enable target latch and print header block
    if ($0 ~ TargetPattern) {
        TargetFound = 1
        TargetLatch = 1
        printf "%sNAME%s\n", SectionColor, ResetColor
        # Add short description if possible
        if (description == "") {
            printf "  %s%s%s - %s\n", TargetColor, Request, ResetColor, Description
        } else {
            printf "  %s%s%s\n", TargetColor, Request, ResetColor
        }
        printf "%sDESCRIPTION%s\n", SectionColor, ResetColor
    }
}

END {
    if (!TargetFound) {
        printf "%sCannot find target `%s`%s\n", ErrorColor, Request, ResetColor
    } else if (!ExtendedDescriptionExists) {
        printf "  %sAn extended description does not exist for this target.%s\n", ErrorColor, ResetColor
    }
}
