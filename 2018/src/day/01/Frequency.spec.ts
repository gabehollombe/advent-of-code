import { Frequency } from './Frequency'

describe('Calculating resulting frequency', () => {
    it('sums sequences of positive integers', () => {
        const input = ['+1', '+1', '+1'];
        expect(new Frequency(input).part1()).toEqual(3);
    });

    it('sums sequences of positive and negative integers', () => {
        const input = ['+1', '+1', '-2'];
        expect(new Frequency(input).part1()).toEqual(0);
    });

    it('sums sequences of negative integers', () => {
        const input = ['-1', '-2', '-3'];
        expect(new Frequency(input).part1()).toEqual(-6);
    });
})

describe('Looping until duplicate frequency', () => {
    it('+1, -1 stops at 0', () => {
        const input = ['+1', '-1'];
        expect(new Frequency(input).part2()).toEqual(0);
    });

    it('+3, +3, +4, -2, -4 stops at 10', () => {
        const input = ['+3', '+3', '+4', '-2', '-4'];
        expect(new Frequency(input).part2()).toEqual(10);
    });

    it('-6, +3, +8, +5, -6 stops at 5', () => {
        const input = ['-6', '+3', '+8', '+5', '-6'];
        expect(new Frequency(input).part2()).toEqual(5);
    });
});