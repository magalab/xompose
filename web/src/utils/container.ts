import { CREATED_STACK, EXITED, RUNNING, UNKNOWN } from "@/types/stack";

export const getContainerExecTerminalName = (
    endpoint: string,
    stackName: string,
    container: string,
    index: number,
) => {
    return "container-exec-" + endpoint + "-" + stackName + "-" + container + "-" + index;
}

/**
 * Possible Inputs:
 * ports:
 *   - "3000"
 *   - "3000-3005"
 *   - "8000:8000"
 *   - "9090-9091:8080-8081"
 *   - "49100:22"
 *   - "8000-9000:80"
 *   - "127.0.0.1:8001:8001"
 *   - "127.0.0.1:5000-5010:5000-5010"
 *   - "6060:6060/udp"
 * @param input
 * @param hostname
 */
export const parseDockerPort = (input: string, hostname: string) => {
    let port;
    let display;

    const parts = input.split("/");
    const part1 = parts[0];
    let protocol = parts[1] || "tcp";

    // Split the last ":"
    const lastColon = part1.lastIndexOf(":");

    if (lastColon === -1) {
        // No colon, so it's just a port or port range
        // Check if it's a port range
        const dash = part1.indexOf("-");
        if (dash === -1) {
            // No dash, so it's just a port
            port = part1;
        } else {
            // Has dash, so it's a port range, use the first port
            port = part1.substring(0, dash);
        }

        display = part1;

    } else {
        // Has colon, so it's a port mapping
        let hostPart = part1.substring(0, lastColon);
        display = hostPart;

        // Check if it's a port range
        const dash = part1.indexOf("-");

        if (dash !== -1) {
            // Has dash, so it's a port range, use the first port
            hostPart = part1.substring(0, dash);
        }

        // Check if it has a ip (ip:port)
        const colon = hostPart.indexOf(":");

        if (colon !== -1) {
            // Has colon, so it's a ip:port
            hostname = hostPart.substring(0, colon);
            port = hostPart.substring(colon + 1);
        } else {
            // No colon, so it's just a port
            port = hostPart;
        }
    }

    let portInt = parseInt(port);

    if (portInt == 443) {
        protocol = "https";
    } else if (protocol === "tcp") {
        protocol = "http";
    }

    return {
        url: protocol + "://" + hostname + ":" + portInt,
        display: display,
    };
}


export const statusConvert = (status: string): number => {
    if (status.startsWith("created")) {
        return CREATED_STACK;
    } else if (status.includes("exited")) {
        return EXITED;
    } else if (status.startsWith("running")) {
        return RUNNING;
    } else {
        return UNKNOWN;
    }
}