export class Coordinate {
    x: number;
    y: number;
    constructor(x: number, y: number) {
        this.x = x;
        this.y = y;
    }

    toString() {
        return `Coordinate(${this.x},${this.y})`;
    }
}

export class Claim {
    id: number
    coordinate: Coordinate
    width: number
    height: number

    constructor(id: number, [x, y]: [number, number], [width, height]: [number, number]) { 
        this.id = id
        this.coordinate = new Coordinate(x, y)
        this.width = width
        this.height = height
    }

    static fromString(string: string) : Claim {
        const parts = string.split(' ')
        const id = parseInt(parts[0].split('#')[1], 10);
        const [x, y] = parts[2].split(':')[0].split(',')
        const [width, height] = parts[3].split('x')
        return new Claim(id, [parseInt(x, 10), parseInt(y, 10)], [parseInt(width, 10), parseInt(height, 10)])
    }
}

function frequencies<T>(list: T[]) : Object {
    return list.reduce(
        (counts, e) => {
            const key = e.toString();
            if (!counts[key]) counts[key] = 0;
            counts[key] += 1;
            return counts;
        }, 
        new Object()
    );
}

export class Fabric {
    private claimedCoordinates: Coordinate[];
    
    constructor() { 
        this.claimedCoordinates = new Array<Coordinate>();
    }

    overlapCount(): any {
        return Object.entries(frequencies(this.claimedCoordinates))
            .map(([_, count]) => count)
            .filter(c => c > 1)
            .length;
    }

    addClaim(claim: Claim) : void {
        let x = claim.coordinate.x;
        let y = claim.coordinate.y;
        let coordinates = Array<Coordinate>();

        while(x < claim.coordinate.x + claim.width) {
            while(y < claim.coordinate.y + claim.height) {
                coordinates.push(new Coordinate(x, y));
                y += 1;
            }
            y = claim.coordinate.y;
            x += 1;
        }

        this.claimedCoordinates = this.claimedCoordinates.concat(coordinates);
    }
}