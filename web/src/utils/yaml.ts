import type { DotenvParseOutput } from "dotenv"
import yaml, { Pair, parseDocument, Scalar, type Document } from "yaml"

// @ts-ignore
import { replaceVariablesSync } from "@inventage/envsubst"

export function envSubstYAML(content: string, env: DotenvParseOutput): string {
    const doc = yaml.parseDocument(content)
    if (doc.contents) {
        // @ts-ignore
        for (const item of doc.contents.items) {
            traverseYAML(item, env)
        }
    }
    return doc.toString()
}

export interface LooseObject {
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    [key: string]: any
}

export const envSubst = (string: string, variables: LooseObject): string => {
    return replaceVariablesSync(string, variables)[0]
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
                let value = item.value as unknown

                if (typeof (value) === "string") {
                    item.value = envSubst(value, env)
                }
            }
        }
        // @ts-ignore
    } else if (pair.value && typeof (pair.value.value) === "string") {
        // @ts-ignore
        pair.value.value = envSubst(pair.value.value, env)
    }
}


/**
 * Copy yaml comments from srcItems to items
 * Attempts to preserve comments by matching content rather than just array indices
 */
// eslint-disable-next-line @typescript-eslint/no-explicit-any
function copyYAMLCommentsItems(items: any, srcItems: any) {
    if (!items || !srcItems) {
        return;
    }

    // First pass - try to match items by their content
    for (let i = 0; i < items.length; i++) {
        // eslint-disable-next-line @typescript-eslint/no-explicit-any
        const item: any = items[i];

        // Try to find matching source item by content
        // eslint-disable-next-line @typescript-eslint/no-explicit-any
        const srcIndex = srcItems.findIndex((srcItem: any) =>
            JSON.stringify(srcItem.value) === JSON.stringify(item.value) &&
            JSON.stringify(srcItem.key) === JSON.stringify(item.key)
        );

        if (srcIndex !== -1) {
            // eslint-disable-next-line @typescript-eslint/no-explicit-any
            const srcItem: any = srcItems[srcIndex];
            // eslint-disable-next-line @typescript-eslint/no-explicit-any
            const nextSrcItem: any = srcItems[srcIndex + 1];

            if (item.key && srcItem.key) {
                item.key.comment = srcItem.key.comment;
                item.key.commentBefore = srcItem.key.commentBefore;
            }

            if (srcItem.comment) {
                item.comment = srcItem.comment;
            }

            // Handle comments between array items
            if (nextSrcItem && nextSrcItem.commentBefore) {
                if (items[i + 1]) {
                    items[i + 1].commentBefore = nextSrcItem.commentBefore;
                }
            }

            // Handle trailing comments after array items
            if (srcItem.value && srcItem.value.comment) {
                if (item.value) {
                    item.value.comment = srcItem.value.comment;
                }
            }

            if (item.value && srcItem.value) {
                if (typeof item.value === "object" && typeof srcItem.value === "object") {
                    item.value.comment = srcItem.value.comment;
                    item.value.commentBefore = srcItem.value.commentBefore;

                    if (item.value.items && srcItem.value.items) {
                        copyYAMLCommentsItems(item.value.items, srcItem.value.items);
                    }
                }
            }
        }
    }
}

export const copyYAMLComments = (doc: Document, src: Document) => {
    doc.comment = src.comment
    doc.commentBefore = src.commentBefore

    if (doc && doc.contents && src && src.contents) {
        // @ts-ignore
        copyYAMLCommentsItems(doc.contents.items, src.contents.items)
    }
}

export const yamlToJSON = (yaml: string) => {
    let doc = parseDocument(yaml);
    if (doc.errors.length > 0) {
        throw doc.errors[0]
    }
    const config = doc.toJS() ?? {}
    if (!config.services) {
        config.services = {}
    }
    if (Array.isArray(config.services) || typeof config.services !== "object") {
        throw new Error("Services must be an object");
    }

    return { config, doc }
}