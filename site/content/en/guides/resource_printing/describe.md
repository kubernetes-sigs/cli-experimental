---
title: "Describe"
linkTitle: "Describe"
weight: 4
type: docs
description: >
    Describes current resources
---


{{< alert color="success" title="TL;DR" >}}
- Print verbose debug information about a Resource
{{< /alert >}}

# Describe Resources

## Motivation

Describe is a **higher level printing operation that may aggregate data from other sources** in addition
to the Resource being queried (e.g. Events).

Describe pulls out the most important information about a Resource from the Resource itself and related
Resources, and formats and prints this information on multiple lines.

- Aggregates data from related Resources
- Formats Verbose Output for debugging

{{< alert color="success" title="Note" >}}
Check out the [References](../../../references) to learn how to print verbose debug information about a Resource
{{< /alert >}}

{{< alert color="warning" title="Get vs Describe" >}}
When Describing a Resource, it may aggregate information from several other Resources.  For instance Describing
a Node will aggregate Pod Resources to print the Node utilization.

When Getting a Resource, it will only print information available from reading that Resource.  While Get may aggregate
data from the *fields* of that Resource, it won't look at fields from other Resources.
{{< /alert >}}