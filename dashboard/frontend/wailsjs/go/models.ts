export namespace platform {
	
	export class SystemStats {
	    id: number;
	    hostname: string;
	    platform: string;
	    cpu: number;
	    ram: number;
	    disk: number;
	    time: number;
	
	    static createFrom(source: any = {}) {
	        return new SystemStats(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.hostname = source["hostname"];
	        this.platform = source["platform"];
	        this.cpu = source["cpu"];
	        this.ram = source["ram"];
	        this.disk = source["disk"];
	        this.time = source["time"];
	    }
	}

}

