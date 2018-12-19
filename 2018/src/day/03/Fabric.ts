export class Coordinate {
    x: number;
    y: number;
    constructor(x: number, y: number) {
        this.x = x;
        this.y = y;
    }

    toString() {
        return `${this.x},${this.y}`;
    }

    static fromString(str: string) {
        const [x,y] = str.split(',')
        return new Coordinate(parseInt(x, 10), parseInt(y, 10))
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

    containsCoordinate(coordinate: Coordinate): boolean {
        if (coordinate.x < this.coordinate.x) return false;
        if (coordinate.y < this.coordinate.y) return false;
        if (coordinate.x >= this.coordinate.x + this.width) return false;
        if (coordinate.y >= this.coordinate.y + this.height) return false;
        return true;
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
    private claims: Claim[];
    
    constructor() { 
        this.claims = new Array<Claim>();
        this.claimedCoordinates = new Array<Coordinate>();
    }

    overlapCount(): any {
        return this.overlappingCoordinates().length;
    }

    private overlappingCoordinates() {
        let freqs = frequencies(this.claimedCoordinates) 
        let overlappingCoordCounts = Object.entries(freqs).filter(([_, count]) => count > 1)
        return overlappingCoordCounts.map(cs => cs[0]);
    }

    private overlappingClaims(): Set<Claim> {
        let claims = new Set<Claim>();
        for (let coordStr of this.overlappingCoordinates()) {
            for (let claim of this.claims.filter(claim => claim.containsCoordinate(Coordinate.fromString(coordStr)))) {
                claims.add(claim);
            }
        }
        return claims
    }

    nonOverlappingClaims(): Set<Claim> {
        const overlappingClaims = this.overlappingClaims()
        this.claims.filter(c => !overlappingClaims.has(c))
        return new Set(
            this.claims.filter(c => !overlappingClaims.has(c))
        )
    }

    addClaim(claim: Claim) : void {
        this.claims.push(claim);

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