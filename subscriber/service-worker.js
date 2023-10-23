importScripts("https://js.pusher.com/beams/service-worker.js");

// curl -H "Content-Type: application/json" \
//      -H "Authorization: Bearer 4ADD87428FB2819A82FFBAD97709CC9AA0DFE2EE0B6E41A320625D5CAF4FA401" \
//      -X POST "https://d6c76406-7cc0-4105-95fc-6329a03549ae.pushnotifications.pusher.com/publish_api/v1/instances/d6c76406-7cc0-4105-95fc-6329a03549ae/publishes" \
//      -d '{"interests":["hello"],"web":{"notification":{"title":"Hello","body":"Hello, world!"}}}'