class GuiController {
    constructor() {
        this.jsonGuiTree = this.getJsonGuiTree()
        this.defaultResponseHandler = null
        this.initPopulateButtons()
    }

    cbPopulateButtons(responseStatus, responseText) {
        let domContent = ""

        const buttonList = JSON.parse(responseText)

        for (const b in buttonList) {
            var butName = buttonList[b]

            domContent += "<div class=\"siimple-btn siimple-btn--orange\" \
            onclick=\"httpGet(null, 'play/" + butName + "',handleServerResponse)\">\
            " + butName + "</div>"
        }

        document.getElementById("audioFilesDisp").innerHTML = domContent
        document.getElementById("respDisp").innerHTML = "Button list retrieved."
    }


    initPopulateButtons() {
        httpGet(this, "getFilelist", this.cbPopulateButtons)
        this.jsonGuiTree.audioFilesBox.innerHTML = "requesting audio files list..."
    }


    attachResponseHandler(handler) {
        this.defaultResponseHandler = handler
    }

    getJsonGuiTree() {
        return {
            queryBox: {
                say: {
                    textInputParam: document.getElementById('param_say')
                },
                mow: {
                    textInputParam: document.getElementById('param_mow')
                },
                guess: {
                    textInputParam: document.getElementById('param_guess')
                },

            },
            queryDisplayBox: {
                textAreaQueryDisp: document.getElementById("queryDisp")
            },
            responseBox: {
                textAreaRespDisp: document.getElementById("respDisp")
            },
            antControlBox: {
                tagDeviceStatus: document.getElementById("deviceStatusTag")
            },
            audioFilesBox: document.getElementById("audioFilesDisp")
        }
    }

    getQueryFromQueryBox(queryType) {
        const prefixApiV2 = 'v2/'
        const isApiV2 = queryType.includes(prefixApiV2);
        queryType = queryType.replace(prefixApiV2, '')
        const query = this.jsonGuiTree.queryBox[queryType].textInputParam.value;
        if (isApiV2) {
            return `${prefixApiV2}${queryType}?query=${query}`;
        } else {
            return `${queryType}/${query}`;
        }
    }


    updateIP(text) {
        this.jsonGuiTree
            .queryBox
            .textInputIpAddr.value = text
    }

    updateQueryDisp(text) {
        this.jsonGuiTree
            .queryDisplayBox
            .textAreaQueryDisp
            .innerHTML = text
    }

    updateRespDisp(text) {
        this.jsonGuiTree
            .responseBox
            .textAreaRespDisp
            .innerHTML = text
    }

    drawDeviceStatus(statusText) {
        const status = {
            color: 'light',
            text: 'Unknown',
        }
        const statusColors = {
            'Online': 'success',
            'Offline': 'dark',
            'Error': 'error',
            'Teapot': 'warning',
        }
        status.color = statusColors[statusText] || status.color
        status.text = statusColors[statusText] || status.text

        if (statusText === 'Teapot') {
            status.text = 'I\'m a teapot!'
        }

        this.jsonGuiTree
            .antControlBox
            .tagDeviceStatus
            .innerHTML = status.text;
        this.jsonGuiTree
            .antControlBox
            .tagDeviceStatus
            .className = `siimple-tag siimple-tag--${status.color} siimple-tag--rounded`;
    }

}