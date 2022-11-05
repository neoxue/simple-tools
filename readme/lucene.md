# what is lucene:

java indexing and search library;



# Terms:

Document

returnable search result item.

Field:

Property, metadata item

Term:

Searchable text

tf/idf:

Term frequency / inverse document frequency



# Documents:

for example:

![filename.doc](https://dzone.com/storage/rc-covers/15329-thumb.png)

### inverted index:

a data structure mapping terms to the documents

indexed fields can be "analyzed", a process of tokenizing and filtering text into individual searchable terms.

![Inverted Index](https://dzone.com/storage/rc-covers/15331-thumb.png)



### Score/lucene practical scoring formula:

​		score(q,d) = coord(q,d) * queryNorm(q) * Sigma[tf(t in d) * idf (t)^2 * t.getBoost() * norm(t,d)]

![formula](https://dzone.com/storage/rc-covers/15333-thumb.png)



| **Factor**       | **Explanation**                                              |
| ---------------- | ------------------------------------------------------------ |
| **score(q,d)**   | The final computed value of numerous factors and weights, numerically representing the relationship between the query and a given document. |
| **coord(q,d)**   | A search-time score factor based on how many of the query terms are found in the specified document. Typically, a document that contains more of the query’s terms will receive a higher score than another document with fewer query terms. |
| **queryNorm(q)** | A normalizing factor used to make scores between queries comparable. This factor does not affect document ranking (since all ranked documents are multiplied by the same factor), but rather just attempts to make scores from different queries (or even different indexes) comparable. |
| **tf(t in d)**   | Correlates to the term’s frequency, defined as the number of times term t appears in the currently scored document d. Documents that have more occurrences of a given term receive a higher score. Note that tf(t in q) is assumed to be 1 and, therefore, does not appear in this equation. However, if a query contains twice the same term, there will be two term-queries with that same term. Hence, the computation would still be correct (although not very efficient). |
| **idf(t)**       | Stands for Inverse Document Frequency. This value correlates to the inverse of docFreq (the number of documents in which the term t appears). This means rarer terms give higher contribution to the total score. idf(t) appears for t in both the query and the document, hence it is squared in the equation. |
| t.getBoost()     | A search-time boost of term t in the query q as specified in the query text (see query syntax), or as set by application calls to setBoost(). |
| norm(t,d)        | Encapsulates a few (indexing time) boost and length factors. |





----



### solr

solr is a application exposing its capabilities thought an easy-to-use http interface.



# Indexing:

### Json:

{

“add”: {
  “doc”: {
        “id”: “doc02”,
        “title”: “Solr JSON”,
        “mimeType”: “application/pdf”}
  }
}

### Comma, or Tab, Separated Values

### Indexing Rich Document Types

### Deleting Documents

POST delete id:""



# Fields

| **Field Attribute** | **Effect and Uses**                                          |
| ------------------- | ------------------------------------------------------------ |
| **stored**          | Stores the original incoming field value in the index. Stored field values are available when documents are retrieved for search results. |
| **term positions**  | Location information of terms within a field. Positional information is necessary for proximity-related queries, such as phrase queries. |
| **term offsets**    | Character begin and end offset values of a term within a fields textual value. Offsets can be handy for increasing performance of generating query term highlighted field fragments. This one typically is a trade-off between highlighting performance and index size. If offsets aren’t stored, they can be computed at highlighting time. |
| **term vectors**    | An “inverted index” structure within a document, containing term/frequency pairs. Term vectors can be useful for more advanced search techniques, such as “more like this” where terms and their frequencies within a single document can be leveraged for finding similar documents. |



# Analysis

charfilter, tokenizer, tokenFilters

Charfilter: ASCII equivalent. for example "Apple/APPLE" -> "apple"  "漢語"   -> "汉语"

tokenizer: "terms"

TokenFilters: add/remove/modify/augment tokens in a sequential pipe line fashion



# Searching

lucene syntax: AND/OR/NOT/+/-



### Debugging Query Parsing

parsedquery:+title:lucene +timestamp:[1266446158657 TO 1297982158657]”

### Explaining Result Scoring

0.8784157 = (MATCH) fieldWeight(title:lucene in 0), product of:
        1.0 = tf(termFreq(title:lucene)=1)
        1.4054651 = idf(docFreq=1, maxDocs=3)
        0.625 = fieldNorm(field=title, doc=0)



# Bells and Whistles











