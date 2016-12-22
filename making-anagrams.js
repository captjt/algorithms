/**
 * Code snippet to solve the strings making anagrams problem
 * Link : https://www.hackerrank.com/challenges/ctci-making-anagrams
 */

process.stdin.resume();
process.stdin.setEncoding('ascii');

var input_stdin = "";
var input_stdin_array = "";
var input_currentline = 0;

process.stdin.on('data', function (data) {
    input_stdin += data;
});

process.stdin.on('end', function () {
    input_stdin_array = input_stdin.split("\n");
    main();
});

function readLine() {
    return input_stdin_array[input_currentline++];
}

/////////////// ignore above this line ////////////////////

function main() {
    var a = Array.from(readLine());
    var b = Array.from(readLine());

    let letterMap = {};
    let deleteCount = 0;

    a.forEach( (l) => {
        if(!letterMap[l]) {
            letterMap[l] = 1;
        } else {
            letterMap[l] += 1;
        }
    });

    b.forEach( (l) => {
        if(!letterMap[l]) {
            deleteCount += 1;
        } else {
            letterMap[l] -= 1;
            if(letterMap[l] < 0) {
                deleteCount += 1;
            }
        }
    })

    for(l in letterMap) {
        if(letterMap[l] > 0) {
            deleteCount += letterMap[l];
        }
    }

    console.log(deleteCount);
}
