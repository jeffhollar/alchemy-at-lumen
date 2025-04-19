// # Copyright 2025 Lunmen Technologies, Inc.
// # All rights reserved.
// #
// # NOTICE:  All information contained herein is, and remains
// # the property of Cisco Systems Incorporated and its suppliers,
// # if any.  The intellectual and technical concepts contained
// # herein are proprietary to Lunmen Technologies Incorporated
// # and its suppliers and may be covered by U.S. and Foreign Patents,
// # patents in process, and are protected by trade secret or copyright law.
// #
// # Dissemination of this information or reproduction of this material
// # is strictly forbidden unless prior written permission is obtained
// # from Cisco Systems Incorporated.
// #

package swagger

import (
	"encoding/json"
	"net/http"
)

func writeJSON(w http.ResponseWriter, status int, data interface{}, wrap string) error {
	wrapper := make(map[string]interface{})
	wrapper[wrap] = data

	js, err := json.Marshal(wrapper)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}
