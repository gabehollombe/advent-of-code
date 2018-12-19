import { Claim, Fabric } from './Fabric';
import { getInputLines } from '../../Input';

const lines = getInputLines(__dirname, 'input.txt')
let fabric = new Fabric()

function parseClaim(line: string): Claim {
    const parts = line.split(' ')
    const id = parseInt(parts[0].split('#')[1], 10);
    const [x, y] = parts[2].split(':')[0].split(',')
    const [width, height] = parts[3].split('x')
    return new Claim(id, [parseInt(x, 10), parseInt(y, 10)], [parseInt(width, 10), parseInt(height, 10)])
}

for (let line of lines) {
    const claim = parseClaim(line)
    fabric.claim(claim.coordinate, [claim.width, claim.height])
}

console.log('Part one: ', fabric.overlapCount())