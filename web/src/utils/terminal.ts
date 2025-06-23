export const TERMINAL_COLS = 105;
export const TERMINAL_ROWS = 10;
export const PROGRESS_TERMINAL_ROWS = 8;
export const COMBINED_TERMINAL_COLS = 58;
export const COMBINED_TERMINAL_ROWS = 20;

export const ERROR_TYPE_VALIDATION = 1;

export const getComposeTerminalName = (endpoint: string, stack: string) => {
    return "compose-" + endpoint + "-" + stack;
}

export const getCombinedTerminalName = (endpoint: string, stack: string) => {
    return "combined-" + endpoint + "-" + stack;
}