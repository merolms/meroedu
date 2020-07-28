
#!/bin/bash

## Init color
red=`tput setaf 1`
green=`tput setaf 2`
yellow=`tput setaf 3`
blue=`tput setaf 4`
orange=`tput setaf 5`
skyblue=`tput setaf 6`
white=`tput setaf 7`
grey=`tput setaf 8`
reset=`tput sgr0`

## Color print
## Uncomment to check which color is associated
# echo "${red}red ${green}green ${yellow}yellow ${blue}blue ${orange}orange ${skyblue}skyblue ${white}white ${grey}grey ${reset}"


## Beautify out with sed output
## Because there is no state value (fail / success / skip) I'm closing with a color
## when a Fail or Success is happen, it will color the next out put.
## Ex: when you got a failure the log print after the failuer will red until
## you have a new test result (Pass/Fail/Skip).

## RUN print is removing because it's print 2 times
beautifyRun="| sed -e 's/\=\=\= RUN\(.*\)//' | grep -v '^$'"

## Pass: {GREN} State {Grey} $Title {RESET} Next
beautifyPass="| sed -e 's/\-\-\- PASS: \(.*\)/${green} ✓ ${grey}\1${reset}/'"
## Pass: {RED State {Grey} $Title {RED} Next
beautifyFail="| sed -e 's/\-\-\- FAIL: \(.*\)/${red} ✖ ${grey}\1${red}/'"
## Pass: {BLUE} State {Grey} $Title {RESET} Next
beautifySkip="| sed -e 's/\-\-\- SKIP: \(.*\)/${blue} ■ ${grey}\1${reset}/'"

## Change the Log format
beautifyLog="| sed -e 's/\t\(.*\).go:\(.*\):\(.*\)/ \t \1.go line:\2\n\t\3/'"

## Change log color line when it's a Curl cmd print
beautifyCurl="| sed -e 's/\t.*.go:.*: curl \(.*\)/${grey} \t ${orange}curl \1${reset}/'"

## beautifyTest $NAME $FILES...
function runTests {
  echo "Run $1 tests..."
  gotest="go test -v "
  cmd="${gotest} ${@:2} ${beautifyRun} ${beautifyPass} ${beautifyFail} ${beautifySkip} ${beautifyCurl} ${beautifyLog}"
  eval $cmd
}

## Main
if [ "$#" -eq 0 ]; then
  echo "  ./beautify.sh [FILE|MODULE]..."
  ## Uncomment and list all of yours modules
  # echo "  Bunch Files:"
  # echo "    - auths"
else
  for var in "$@"
  do
    echo "$var"
    case $var in
      *.go)
        runTests $1 $1.go
      ;;
      ## Exemple of module
      # auths)
      #   files="before_test.go"
      #   files+="authentifications_register_test.go "
      #   files+="authentifications_login_test.go "
      #   files+="authentifications_logout_test.go"
      #   beautifyTest "authentifications" $files
      # ;;
    esac
  done
fi
