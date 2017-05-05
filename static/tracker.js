// singleton def
var Tracker = new function() {
  // properties
  var _self = this;
  var url = null;

  // public initialisation method
  _self.init = function() {
    console.log(this)
    setUrl(window.location.origin);
    addResizeListener();
  };
  
  // private gets/sets
  var setUrl = function(currentUrl) {
    // check url and throw
    url = currentUrl;
  };

  var getUrl = function() {
    return url;
  };

  var setSessionId = function(sessionId) {
    if (typeof(Storage) !== "undefined") {
      localStorage.setItem("ravelin_sessionId", sessionId);
    } else {
        // No web storage support. To implement revert back to cookie
    }
  };

  var getSessionId = function() {
    return localStorage.getItem("ravelin_sessionId");
  };

  var addCopyPasteListener = function() {};

  var addResizeListener = function() {
    window.addEventListener("resize", handleResize);
  }
  var handleResize = function(event) {
    console.log("resize", event);
    window.removeEventListener("resize", handleResize);
    // postData
  }

};


$(document).ready(function() {
  Tracker.init();
});