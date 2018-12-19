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

    claim(claimTopLeft: Coordinate, [width, height]) : void {
        let x = claimTopLeft.x;
        let y = claimTopLeft.y;
        let coordinates = Array<Coordinate>();
        while(x < claimTopLeft.x + width) {
            while(y < claimTopLeft.y + height) {
                coordinates.push(new Coordinate(x, y));
                y += 1;
            }
            y = claimTopLeft.y;
            x += 1;
        }
        this.claimedCoordinates = this.claimedCoordinates.concat(coordinates);
    }
}