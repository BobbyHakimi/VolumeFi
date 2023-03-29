# VolumeFi
Determine a trip's itinerary from a list of airport codes

# API Specifications
/calculate -
Takes an array of strings arrays (in JSON) representing the unordered complete trip itinerary. Returns a string slice representing the starting and ending airports.

Example:
[["IND", "EWR"], ["SFO", "ATL"], ["GSO", "IND"], ["ATL", "GSO"]] => ["SFO", "EWR"]