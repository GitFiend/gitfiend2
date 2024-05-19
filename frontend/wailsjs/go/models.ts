export namespace main {
	
	export class MyLoot {
	    gems: number;
	    bottleCaps: number;
	
	    static createFrom(source: any = {}) {
	        return new MyLoot(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.gems = source["gems"];
	        this.bottleCaps = source["bottleCaps"];
	    }
	}

}

