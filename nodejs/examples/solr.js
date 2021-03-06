const solr = require('solr-node');
const config = require("platformsh-config").config();

exports.usageExample = async function() {

    let client = new solr(config.formattedCredentials('solr', 'solr-node'));

    let output = '';

    // Add a document.
    var result = await client.update({
        id: 123,
        name: 'Valentina Tereshkova',
    });

    output += "Adding one document. Status (0 is success): " + result.responseHeader.status +  "<br />\n";

    // Flush writes so that we can query against them.
    await client.softCommit();

    // Select one document:
    let strQuery = client.query().q();
    result = await client.search(strQuery);
    output += "Selecting documents (1 expected): " + result.response.numFound + "<br />\n";

    // Delete one document.
    result = await client.delete({id: 123});
    output += "Deleting one document. Status (0 is success): " + result.responseHeader.status + "<br />\n";

    return output;
};
