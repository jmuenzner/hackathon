package Legacy

import (
  "github.com/pkg/errors"
  "github.com/autopilothq/lg"
  ct "github.com/autopilothq/banks/contract/types"
  proto "github.com/autopilothq/banks/protocol"

  // This is the protocol definitions for journeys itself.
  "github.com/hackathon/journeys/resources/Journeys/protocol"

  "github.com/hackathon/journeys/dataAccess/couch"
)

var (
  // ErrReadFailed indicates that the result from Cloudant was either not valid
  // JSON or was otherwise badly formed.
  ErrCloudantFailed = "Something with Cloudant went boom!"
)

type Doc struct {
  Ids []string `json:"childIds"`
}

// GetTree is the actual function that will be executed by Banks core every
// time a `GetTree` message is dispatched to `Journeys`.
func GetTree(request *proto.Request, log lg.Log, aux ct.Auxiliary) *proto.Response {

  // Extract out the request body (GetTreeRequest) for GetTree.
  // An empty GetTreeRequest was generated for you by "banks apply", it lives
  // inside resources/Journeys/protocol/protocol.proto.
  getTreeRequest := request.Body.(*protocol.GetTreeRequest)
  log.Infof("REQUEST %v\n", getTreeRequest)

  db, err := couch.Db("bislr_development")
  if err != nil {
    // This code indicates to the client that we attempted to read from a source
    // external to the Banks process (indicating IPC or network traversal) but
    // that read failed.
    code := proto.ResponseCode_ERR_READ_FAILED

    // MakeFatalResponse is like MakeOkResponse in that it generates an empty
    // response. It differs in that it generates a failed response.
    return request.MakeFatalResponse(code,
      errors.Wrap(err, ErrCloudantFailed).Error())
  }

  var doc Doc
  err = db.GetDocument("journeys_tree", &doc, nil)
  if err != nil {
    // This code indicates to the client that we attempted to read from a source
    // external to the Banks process (indicating IPC or network traversal) but
    // that read failed.
    code := proto.ResponseCode_ERR_READ_FAILED

    // MakeFatalResponse is like MakeOkResponse in that it generates an empty
    // response. It differs in that it generates a failed response.
    return request.MakeFatalResponse(code,
      errors.Wrap(err, ErrCloudantFailed).Error())
  }
  log.Infof("RESULT DATA %v\n", doc.Ids)

  // creates a generic success response, this would be the equivilant to
  // something like HTTP Status 200.
  response := request.MakeOkResponse()
  body := response.Body.(*protocol.GetTreeResponse)
  body.ChildIds = doc.Ids

  // Lastly, the handler returns the final response. This will be relayed
  // back to the original client by Banks core.
  return response
}