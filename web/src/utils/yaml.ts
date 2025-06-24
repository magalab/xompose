import type { DotenvParseOutput } from "dotenv";
import yaml, { Pair, parseDocument, Scalar, type Document } from "yaml";

// @ts-ignore
import { replaceVariablesSync } from "@inventage/envsubst";

export function envSubstYAML(content: string, env: DotenvParseOutput): string {
    const doc = yaml.parseDocument(content);
    if (doc.contents) {
        // @ts-ignore
        for (const item of doc.contents.items) {
            traverseYAML(item, env);
        }
    }
    return doc.toString();
}

export interface LooseObject {
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    [key: string]: any
}

export const envSubst = (string: string, variables: LooseObject): string => {
    return replaceVariablesSync(string, variables)[0];
}

/**
 * Used for envsubstYAML(...)
 * @param pair
 * @param env
 */
const traverseYAML = (pair: Pair, env: DotenvParseOutput): void => {
    // @ts-ignore
    if (pair.value && pair.value.items) {
        // @ts-ignore
        for (const item of pair.value.items) {
            if (item instanceof Pair) {
                traverseYAML(item, env);
            } else if (item instanceof Scalar) {
                let value = item.value as unknown;

                if (typeof (value) === "string") {
                    item.value = envSubst(value, env);
                }
            }
        }
        // @ts-ignore
    } else if (pair.value && typeof (pair.value.value) === "string") {
        // @ts-ignore
        pair.value.value = envSubst(pair.value.value, env);
    }
}


export const copyYAMLComments = (doc: Document, src: Document) => {
    doc.comment = src.comment;
    doc.commentBefore = src.commentBefore;

    if (doc && doc.contents && src && src.contents) {
        // @ts-ignore
        copyYAMLCommentsItems(doc.contents.items, src.contents.items);
    }
}

export const yamlToJSON = (yaml: string) => {
    let doc = parseDocument(yaml);
    if (doc.errors.length > 0) {
        throw doc.errors[0];
    }

    const config = doc.toJS() ?? {};

    // Check data types
    // "services" must be an object
    if (!config.services) {
        config.services = {};
    }

    if (Array.isArray(config.services) || typeof config.services !== "object") {
        throw new Error("Services must be an object");
    }

    return { config, doc }
}