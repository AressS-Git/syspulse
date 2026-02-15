export namespace platform {
	
	export class SystemStats {
	    id: number;
	    hostname: string;
	    platform: string;
	    cpu: number;
	    ram: number;
	    disk: number;
	    incoming_net_traffic: number;
	    outbound_net_traffic: number;
	    processes: string[];
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
	        this.incoming_net_traffic = source["incoming_net_traffic"];
	        this.outbound_net_traffic = source["outbound_net_traffic"];
	        this.processes = source["processes"];
	        this.time = source["time"];
	    }
	}

}

