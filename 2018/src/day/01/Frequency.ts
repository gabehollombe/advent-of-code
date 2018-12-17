export class Frequency {
    private numbers: number[];

    constructor(input: string[]) {
        this.numbers = input.map(n => parseInt(n, 10));
    }

    part1(): number {
        return this.numbers
            .reduce(((sum, i) => sum + i), 0);
    }

    part2(): number {
        const length = this.numbers.length;

        let seen = {0: true};
        let index = 0;
        let sum = 0;
        while (true) {
            sum = sum + this.numbers[index % length];
            if (seen[sum]) return sum;
            seen[sum] = true;
            index = index + 1;
        }
    }
}